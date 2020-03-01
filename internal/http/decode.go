package http

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/denniszl/wallet_flexing/internal/endpoints"
	httptransport "github.com/go-kit/kit/transport/http"
)

// Decoder type to be for things like wrappers -- similar to endpoint.Endpoint.
type Decoder func(c context.Context, r *http.Request) (request interface{}, err error)

// DecoderWrapper is the return type of functions that will be used for chaining
// Decoder code.
type DecoderWrapper func(Decoder) httptransport.DecodeRequestFunc

// decodeRequest is a default decoder when we don't need anything special
func decodeRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	defer r.Body.Close()
	return nil, nil
}

func decodeGetTransactions(_ context.Context, r *http.Request) (request interface{}, err error) {
	queryParams := r.URL.Query()
	// we don't read it though
	defer r.Body.Close()
	var startTime time.Time
	var endTime time.Time
	startDateTime, ok1 := queryParams["start_date_time"]
	if ok1 && len(startDateTime) > 0 {
		startTime, err = time.Parse(time.RFC3339, startDateTime[0])
		if err != nil {
			return nil, ErrInvalidTimestampsFormat
		}
	} else {
		ok1 = false
	}

	endDateTime, ok2 := queryParams["end_date_time"]
	if ok2 && len(endDateTime) > 0 {
		endTime, err = time.Parse(time.RFC3339, endDateTime[0])
		if err != nil {
			return nil, ErrInvalidTimestampsFormat
		}
	} else {
		ok2 = false
	}

	req := endpoints.GetTransactionsRequest{}

	if ok1 && ok2 {
		fmt.Println(startTime, endTime)
		if !startTime.Before(endTime) {
			return nil, ErrInvalidTimestamps
		}
		req.From = startTime
		req.To = endTime
	} else if ok1 && !ok2 {
		return nil, ErrMissingTimestamp
	} else if !ok1 && ok2 {
		return nil, ErrMissingTimestamp
	} else {
		req.GetAll = true
	}
	return req, nil
}

func decodePostTransaction(_ context.Context, r *http.Request) (request interface{}, err error) {
	requestObj := endpoints.PostTransactionRequest{}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &requestObj) // Purposely ignore error here
	if err != nil {
		return nil, ErrInvalidBody
	}

	return requestObj, nil
}
