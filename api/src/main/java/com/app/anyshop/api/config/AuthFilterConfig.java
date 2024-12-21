//package com.app.anyshop.api.config;
//
//import static org.springframework.cloud.gateway.support.ServerWebExchangeUtils.GATEWAY_ORIGINAL_REQUEST_URL_ATTR;
//import static org.springframework.cloud.gateway.support.ServerWebExchangeUtils.GATEWAY_REQUEST_URL_ATTR;
//import static org.springframework.cloud.gateway.support.ServerWebExchangeUtils.GATEWAY_ROUTE_ATTR;
//
//import com.app.anyshop.api.viewmodel.AuthValidateVm;
//import com.app.anyshop.api.viewmodel.ResVm;
//import com.fasterxml.jackson.core.JsonProcessingException;
//import com.fasterxml.jackson.databind.ObjectMapper;
//import java.net.URI;
//import java.util.Collections;
//import java.util.Date;
//import java.util.Objects;
//import java.util.Set;
//import org.slf4j.Logger;
//import org.slf4j.LoggerFactory;
//import org.springframework.beans.factory.annotation.Autowired;
//import org.springframework.beans.factory.annotation.Value;
//import org.springframework.cloud.gateway.filter.GatewayFilter;
//import org.springframework.cloud.gateway.filter.factory.AbstractGatewayFilterFactory;
//import org.springframework.cloud.gateway.route.Route;
//import org.springframework.core.io.buffer.DataBufferFactory;
//import org.springframework.http.HttpMethod;
//import org.springframework.http.HttpStatus;
//import org.springframework.http.HttpStatusCode;
//import org.springframework.http.server.reactive.ServerHttpRequest;
//import org.springframework.http.server.reactive.ServerHttpResponse;
//import org.springframework.stereotype.Component;
//import org.springframework.web.reactive.function.client.WebClient;
//import org.springframework.web.reactive.function.client.WebClientResponseException;
//import org.springframework.web.server.ServerWebExchange;
//import reactor.core.publisher.Mono;
//
//@Component
//public class AuthFilterConfig extends AbstractGatewayFilterFactory<AuthFilterConfig.Config> {
//
//  private static final Logger LOG = LoggerFactory.getLogger(AuthFilterConfig.class);
//  private static final String IAM_SERVICE = "iam";
//  private static final String X_API_SECRET_TOKEN_KEY = "X-API-SECRET-TOKEN";
//  private static final String X_API_USER_ID = "X-API-USER-ID";
//  private static final String X_API_USER = "X-API-USER";
//  private static final String X_API_USER_ROLE = "X-API-USER-ROLE";
//  private final WebClient.Builder webClientBuilder;
//
//  @Autowired private ObjectMapper objectMapper;
//
//  @Value("${host.iam}")
//  private String AUTH_URI;
//
//  public AuthFilterConfig(WebClient.Builder webClientBuilder) {
//    super(Config.class);
//    this.webClientBuilder = webClientBuilder;
//  }
//
//  @Override
//  public GatewayFilter apply(Config config) {
//    String validateUrl = AUTH_URI + "/api/v1/validate";
//    return (exchange, chain) -> {
//      ServerHttpRequest request = exchange.getRequest();
//      String path = request.getURI().getPath();
//      HttpMethod method = request.getMethod();
//      LOG.info("*****************************************************************************");
//      Set<URI> uris =
//          exchange.getAttributeOrDefault(GATEWAY_ORIGINAL_REQUEST_URL_ATTR, Collections.emptySet());
//      String originalUri = uris.isEmpty() ? "Unknown" : uris.iterator().next().toString();
//
//      // Also exclude all the incoming request to IAM service.
//      if (originalUri.contains(IAM_SERVICE)) {
//        return chain.filter(exchange);
//      }
//
//      Route route = exchange.getAttribute(GATEWAY_ROUTE_ATTR);
//      URI routeUri = exchange.getAttribute(GATEWAY_REQUEST_URL_ATTR);
//      if (route == null || routeUri == null) {
//        return chain.filter(exchange);
//      }
//      LOG.info(
//          "Incoming request %s is routed to id: %s, uri: %s"
//              .formatted(originalUri, route.getId(), routeUri));
//
//      // TODO: change get token from header instead of `authorization`
//      if (ExcludeUrlConfig.isSecure(path, method)) {
//        String secretToken =
//            Objects.requireNonNull(request.getHeaders().get(X_API_SECRET_TOKEN_KEY)).get(0);
//        String userId = Objects.requireNonNull(request.getHeaders().get(X_API_USER_ID)).get(0);
//        if (secretToken == null || userId == null) {
//          LOG.error("Can not access to secure end point with null token");
//          return onError(
//              exchange,
//              "Can not access to secure end point " + "with null token",
//              "Token is null",
//              HttpStatus.UNAUTHORIZED);
//        }
//        return webClientBuilder
//            .build()
//            .get()
//            .uri(validateUrl)
//            .header(X_API_SECRET_TOKEN_KEY, secretToken)
//            .header(X_API_USER_ID, userId)
//            .retrieve()
//            .bodyToMono(AuthValidateVm.class)
//            .map(
//                res -> {
//                  ServerHttpRequest mutatedRequest =
//                      exchange
//                          .getRequest()
//                          .mutate()
//                          .header(X_API_USER, res.username())
//                          .header(X_API_USER_ROLE, res.role())
//                          .header(X_API_USER_ID, res.userId())
//                          .build();
//                  LOG.info(
//                      "Success validate user. Forward to %s"
//                          .formatted(exchange.getRequest().getPath()));
//
//                  return exchange.mutate().request(mutatedRequest).build();
//                })
//            .flatMap(chain::filter)
//            .onErrorResume(
//                err -> {
//                  LOG.error("Error when validating account");
//                  HttpStatusCode errCode = null;
//                  String errMsg = "";
//                  if (err instanceof WebClientResponseException exception) {
//                    errCode = exception.getStatusCode();
//                    errMsg = exception.getMessage();
//                  } else {
//                    errCode = HttpStatus.BAD_GATEWAY;
//                    errMsg = HttpStatus.BAD_GATEWAY.getReasonPhrase();
//                  }
//                  return onError(exchange, errMsg, "JWT Authentication Failed", errCode);
//                });
//      }
//      LOG.info("*****************************************************************************");
//      return chain.filter(exchange);
//    };
//  }
//
//  private Mono<Void> onError(
//      ServerWebExchange exchange, String err, String errDetails, HttpStatusCode httpStatus) {
//    DataBufferFactory dataBufferFactory = exchange.getResponse().bufferFactory();
//    ServerHttpResponse response = exchange.getResponse();
//    response.setStatusCode(httpStatus);
//    try {
//      response.getHeaders().add("Content-Type", "application/json");
//      ResVm data = new ResVm(httpStatus, err, errDetails, null, new Date());
//      byte[] byteData = objectMapper.writeValueAsBytes(data);
//      return response.writeWith(Mono.just(byteData).map(dataBufferFactory::wrap));
//
//    } catch (JsonProcessingException e) {
//      LOG.error(e.getMessage());
//    }
//    return response.setComplete();
//  }
//
//  public static class Config {
//
//    public Config() {
//      // TODO document why this constructor is empty
//    }
//  }
//}
