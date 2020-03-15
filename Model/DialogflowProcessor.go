package Model

import (
	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"context"
	"fmt"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"google.golang.org/api/option"
	dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"log"
	"strconv"
)

// DialogflowProcessor has all the information for connecting with Dialogflow
type DialogflowProcessor struct {
	projectID        string
	authJSONFilePath string
	language         string
	timeZone         string
	sessionClient    *dialogflow.SessionsClient
	context          context.Context
}

func (dp *DialogflowProcessor) Init(input ...string) (err error) {
	dp.projectID = input[0]
	dp.authJSONFilePath = input[1]
	dp.language = input[2]
	dp.timeZone = input[3]
	dp.context = context.Background()

	// Creation of session with credentials to DialogFlow
	sessionClient, err := dialogflow.NewSessionsClient(dp.context, option.WithCredentialsFile(dp.authJSONFilePath))
	if err != nil {
		log.Fatal("Auth error with Dialogflow")
	}
	dp.sessionClient = sessionClient

	// Returning the object (DialogFlowProcessor) itself
	return
}

func (dp *DialogflowProcessor) ProcessNLP(rawMessage string, username string) (nlpResponse NLPResponse) {
	// The sessionID is the username of our client
	sessionID := username

	// Request for each session
	request := dialogflowpb.DetectIntentRequest{
		Session: fmt.Sprintf("projects/%s/agent/sessions/%s", dp.projectID, sessionID),
		QueryInput: &dialogflowpb.QueryInput{
			Input: &dialogflowpb.QueryInput_Text{
				Text: &dialogflowpb.TextInput{
					Text:         rawMessage,
					LanguageCode: dp.language,
				},
			},
		},
		QueryParams: &dialogflowpb.QueryParameters{
			TimeZone: dp.timeZone,
		},
	}

	// Receive DialogFlowResponse
	response, err := dp.sessionClient.DetectIntent(dp.context, &request)
	if err != nil {
		log.Fatalf("Error in communication with Dialogflow %s", err.Error())
		return
	}
	queryResult := response.GetQueryResult()

	// Save on response the Intent, Confidence of our response and response message
	if queryResult.Intent != nil {
		nlpResponse.Intent = queryResult.Intent.DisplayName
		nlpResponse.Confidence = queryResult.IntentDetectionConfidence
		nlpResponse.ResponseMessage = queryResult.FulfillmentText
	}

	// Mapping the entities that DialogFlow detected, if any
	nlpResponse.Entities = make(map[string]string)
	params := queryResult.Parameters.GetFields()
	if len(params) > 0 {
		for paramName, p := range params {
			fmt.Printf("Param %s: %s (%s)", paramName, p.GetStringValue(), p.String())
			extractedValue := extractDialogFlowEntities(p)
			nlpResponse.Entities[paramName] = extractedValue
		}
	}

	// Returning the object (DialogFlowProcessor) itself
	return
}

func extractDialogFlowEntities(p *structpb.Value) (extractedEntity string) {
	//Function to extract the entities retrieved by DialogFlow
	// Next, we define the different types that we are supporting

	kind := p.GetKind()
	switch kind.(type) {
	case *structpb.Value_StringValue:
		return p.GetStringValue()
	case *structpb.Value_NumberValue:
		return strconv.FormatFloat(p.GetNumberValue(), 'f', 6, 64)
	case *structpb.Value_BoolValue:
		return strconv.FormatBool(p.GetBoolValue())
	case *structpb.Value_StructValue:
		s := p.GetStructValue()
		fields := s.GetFields()
		extractedEntity = ""
		for key, value := range fields {
			if key == "amount" {
				extractedEntity = fmt.Sprintf("%s%s", extractedEntity, strconv.FormatFloat(value.GetNumberValue(), 'f', 6, 64))
			}
			if key == "unit" {
				extractedEntity = fmt.Sprintf("%s%s", extractedEntity, value.GetStringValue())
			}
			if key == "date_time" {
				extractedEntity = fmt.Sprintf("%s%s", extractedEntity, value.GetStringValue())
			}
			// @TODO: Other entity types can be added here
		}
		return extractedEntity
	case *structpb.Value_ListValue:
		list := p.GetListValue()
		if len(list.GetValues()) > 1 {
			// @TODO: Extract more values
		}
		extractedEntity = extractDialogFlowEntities(list.GetValues()[0])
		return extractedEntity
	default:
		return ""
	}
}
