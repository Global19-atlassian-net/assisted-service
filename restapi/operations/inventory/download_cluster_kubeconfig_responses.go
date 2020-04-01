// Code generated by go-swagger; DO NOT EDIT.

package inventory

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"
)

// DownloadClusterKubeconfigOKCode is the HTTP code returned for type DownloadClusterKubeconfigOK
const DownloadClusterKubeconfigOKCode int = 200

/*DownloadClusterKubeconfigOK The kubeconfig file

swagger:response downloadClusterKubeconfigOK
*/
type DownloadClusterKubeconfigOK struct {

	/*
	  In: Body
	*/
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewDownloadClusterKubeconfigOK creates DownloadClusterKubeconfigOK with default headers values
func NewDownloadClusterKubeconfigOK() *DownloadClusterKubeconfigOK {

	return &DownloadClusterKubeconfigOK{}
}

// WithPayload adds the payload to the download cluster kubeconfig o k response
func (o *DownloadClusterKubeconfigOK) WithPayload(payload io.ReadCloser) *DownloadClusterKubeconfigOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the download cluster kubeconfig o k response
func (o *DownloadClusterKubeconfigOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DownloadClusterKubeconfigOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// DownloadClusterKubeconfigNotFoundCode is the HTTP code returned for type DownloadClusterKubeconfigNotFound
const DownloadClusterKubeconfigNotFoundCode int = 404

/*DownloadClusterKubeconfigNotFound Cluster not found

swagger:response downloadClusterKubeconfigNotFound
*/
type DownloadClusterKubeconfigNotFound struct {
}

// NewDownloadClusterKubeconfigNotFound creates DownloadClusterKubeconfigNotFound with default headers values
func NewDownloadClusterKubeconfigNotFound() *DownloadClusterKubeconfigNotFound {

	return &DownloadClusterKubeconfigNotFound{}
}

// WriteResponse to the client
func (o *DownloadClusterKubeconfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// DownloadClusterKubeconfigInternalServerErrorCode is the HTTP code returned for type DownloadClusterKubeconfigInternalServerError
const DownloadClusterKubeconfigInternalServerErrorCode int = 500

/*DownloadClusterKubeconfigInternalServerError Internal server error

swagger:response downloadClusterKubeconfigInternalServerError
*/
type DownloadClusterKubeconfigInternalServerError struct {
}

// NewDownloadClusterKubeconfigInternalServerError creates DownloadClusterKubeconfigInternalServerError with default headers values
func NewDownloadClusterKubeconfigInternalServerError() *DownloadClusterKubeconfigInternalServerError {

	return &DownloadClusterKubeconfigInternalServerError{}
}

// WriteResponse to the client
func (o *DownloadClusterKubeconfigInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}