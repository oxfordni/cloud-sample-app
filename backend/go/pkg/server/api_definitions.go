package server

import (
	swagger "github.com/davidebianchi/gswagger"
)

var withQuoteId = swagger.ParameterValue{
	"id": {
		Schema:      &swagger.Schema{Value: 0},
		Description: "id is a quote id number above 0",
	},
}

var oneQuoteContent = swagger.ContentValue{
	Content: swagger.Content{
		"application/json": {Value: Quote{}},
	},
}

var oneQuoteResponse = map[int]swagger.ContentValue{
	200: oneQuoteContent,
}

var HealthHandlerDefinition = swagger.Definitions{
	Responses: map[int]swagger.ContentValue{
		200: {
			Content: swagger.Content{
				"application/json": {Value: HealthStatus{}},
			},
		},
	},
}

var MovieQuotesHandlerDefinition = swagger.Definitions{
	Responses: oneQuoteResponse,
}

var PostQuoteHandlerDefinition = swagger.Definitions{
  RequestBody: &oneQuoteContent,
	Responses: oneQuoteResponse,
}

var GetAllQuoteHandlerDefinition = swagger.Definitions{
	Responses: map[int]swagger.ContentValue{
		200: {
			Content: swagger.Content{
				"application/json": {Value: []Quote{}},
			},
		},
	},
}

var GetQuoteHandlerDefinition = swagger.Definitions{
	PathParams: withQuoteId,
	Responses: oneQuoteResponse,
}

var PutQuoteHandlerDefinition = swagger.Definitions{
	PathParams: withQuoteId,
  RequestBody: &oneQuoteContent,
	Responses: oneQuoteResponse,
}

var DeleteQuoteHandlerDefinition = swagger.Definitions{
	PathParams: withQuoteId,
	Responses: map[int]swagger.ContentValue{
    201: {
      Content: swagger.Content{
        "text/html": {Value: ""},
      },
    },
	},
}
