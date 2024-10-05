1xx - INFORMATIONAL RESONSE

100. http.StatusContinue: The server has received the request headers and the client should proceed to send the request body.
101. http.StatusSwitchingProtocols: The server understands and is willing to comply with the client's request, via the Upgrade header field, for a change in the application protocol being used on this connection.


2xx - SUCCESS

200. http.StatusOK: The request has succeeded.
201. http.StatusCreated: The request has been fulfilled and resulted in a new resource being created.
202. http.StatusAccepted: The request has been accepted for processing, but the processing has not been completed.
203. http.StatusNonAuthoritativeInfo: The server successfully processed the request, but is returning information that may be from another source.
204. http.StatusNoContent: The server successfully processed the request, but is not returning any content.
205. http.StatusResetContent: The server successfully processed the request, but is not returning any content. Unlike a 204 response, this response requires that the requester reset the document view.
206. http.StatusPartialContent: The server is delivering only part of the resource due to a range header sent by the client.


3xx - REDIRECTIONS

300. http.StatusMultipleChoices: The target resource has more than one representation, each with its own more specific identifier, and information about the alternatives is being provided so that the user (or user agent) can select a preferred representation and redirect its request to that location.
301. http.StatusMovedPermanently: The target resource has been assigned a new permanent URI, and any future references to this resource ought to use one of the enclosed URIs.
302. http.StatusFound: The target resource resides temporarily under a different URI.
303. http.StatusSeeOther: The server is redirecting the user agent to a different resource, as indicated by a URI in the Location header field, which may or may not be a different URI than the one the client originally requested.
304. http.StatusNotModified: The server has fulfilled the request but does not need to return an entity-body, and might want to return updated metainformation.
305. http.StatusUseProxy: Many HTTP clients (such as Mozilla and Internet Explorer) do not correctly handle responses with this status code.
307. http.StatusTemporaryRedirect: In this case, the request should be repeated with another URI; however, future requests can still be directed to the original URI.
308. http.StatusPermanentRedirect: This and all future requests should be directed to the given URI.


4xx - CLIENT ERRORS

400. http.StatusBadRequest: The request could not be understood by the server due to malformed syntax.
401. http.StatusUnauthorized: The request has not been applied because it lacks valid authentication credentials for the target resource.
402. http.StatusPaymentRequired: Reserved for future use.
403. http.StatusForbidden: The server understood the request, but is refusing to fulfill it.
404. http.StatusNotFound: The origin server did not find a current representation for the target resource or is not willing to disclose that one exists.
405. http.StatusMethodNotAllowed: The method received in the request-line is known by the origin server but not supported by the target resource.
406. http.StatusNotAcceptable: The target resource does not have a current representation that would be acceptable to the user agent, according to the proactive negotiation header fields received in the request.
407. http.StatusProxyAuthRequired: Similar to 401 Unauthorized, but it indicates that the client needs to authenticate itself in order to use a proxy.
408. http.StatusRequestTimeout: The server would like to shut down this unused connection.
409. http.StatusConflict: Indicates that the request could not be processed because of conflict in the request, such as an edit conflict.
410. http.StatusGone: Indicates that the resource requested is no longer available and will not be available again.
411. http.StatusLengthRequired: The request did not specify the length of its content, which is required by the requested resource.
412. http.StatusPreconditionFailed: The server does not meet one of the preconditions that the requester put on the request.
413. http.StatusRequestEntityTooLarge: The request is larger than the server is willing or able to process.
`http