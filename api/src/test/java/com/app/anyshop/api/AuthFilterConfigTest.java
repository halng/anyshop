//package com.app.anyshop.api;
//
//import static org.mockito.ArgumentMatchers.*;
//import static org.mockito.Mockito.*;
//import static org.springframework.cloud.gateway.support.ServerWebExchangeUtils.GATEWAY_REQUEST_URL_ATTR;
//import static org.springframework.cloud.gateway.support.ServerWebExchangeUtils.GATEWAY_ROUTE_ATTR;
//
//import com.app.anyshop.api.config.AuthFilterConfig;
//import java.net.URI;
//import java.net.URISyntaxException;
//import java.util.Set;
//import org.junit.jupiter.api.Assertions;
//import org.junit.jupiter.api.BeforeEach;
//import org.junit.jupiter.api.Test;
//import org.junit.jupiter.api.extension.ExtendWith;
//import org.mockito.InjectMocks;
//import org.mockito.Mock;
//import org.mockito.junit.jupiter.MockitoExtension;
//import org.springframework.cloud.gateway.filter.GatewayFilter;
//import org.springframework.cloud.gateway.filter.GatewayFilterChain;
//import org.springframework.cloud.gateway.route.Route;
//import org.springframework.http.HttpHeaders;
//import org.springframework.http.HttpMethod;
//import org.springframework.http.server.reactive.ServerHttpRequest;
//import org.springframework.http.server.reactive.ServerHttpResponse;
//import org.springframework.web.reactive.function.client.WebClient;
//import org.springframework.web.server.ServerWebExchange;
//import reactor.core.publisher.Mono;
//
//@ExtendWith(MockitoExtension.class)
//public class AuthFilterConfigTest {
//
//  @InjectMocks private AuthFilterConfig config;
//
//  @Mock private WebClient.Builder webClientBuilder;
//
//  @Mock private WebClient webClient;
//
//  @Mock private WebClient.RequestHeadersUriSpec<?> requestHeadersUriSpec;
//
//  @Mock private WebClient.ResponseSpec responseSpec;
//
//  @Mock private ServerWebExchange exchange;
//
//  @Mock private ServerHttpRequest request;
//
//  @Mock private ServerHttpResponse response;
//
//  @Mock private GatewayFilterChain chain;
//
//  @Mock private Route route;
//
//  @BeforeEach
//  void setUp() {
//    lenient().when(webClientBuilder.build()).thenReturn(webClient);
//    lenient().when(exchange.getRequest()).thenReturn(request);
//    lenient().when(exchange.getResponse()).thenReturn(response);
//  }
//
//  @Test
//  void testGatewayWithoutValidateRequest() throws URISyntaxException {
//    var simulateRequestURI = "http://localhost/login";
//    var requestUri = new URI(simulateRequestURI);
//    when(request.getURI()).thenReturn(new URI(simulateRequestURI));
//    when(request.getMethod()).thenReturn(HttpMethod.POST);
//
//    when(exchange.getAttributeOrDefault(anyString(), anySet()))
//        .thenReturn(Set.of(new URI("login")));
//    when(exchange.getAttribute(GATEWAY_ROUTE_ATTR)).thenReturn(route);
//    when(exchange.getAttribute(GATEWAY_REQUEST_URL_ATTR)).thenReturn(requestUri);
//
//    GatewayFilter filter = config.apply(new AuthFilterConfig.Config());
//    Mono<Void> result = filter.filter(exchange, chain);
//
//    Assertions.assertNull(result);
//  }
//
//    @Test
//  void testGatewayWhenIncomingToIAM() throws URISyntaxException {
//    var simulateRequestURI = "http://localhost/iam";
//    when(request.getURI()).thenReturn(new URI(simulateRequestURI));
//    when(request.getMethod()).thenReturn(HttpMethod.POST);
//
//    when(exchange.getAttributeOrDefault(anyString(), anySet()))
//        .thenReturn(Set.of(new URI("login")));
//
//    GatewayFilter filter = config.apply(new AuthFilterConfig.Config());
//    Mono<Void> result = filter.filter(exchange, chain);
//
//    Assertions.assertNull(result);
//  }
//
//   @Test
//  void testGatewayWithoutAuthorizationInHeader() throws URISyntaxException {
//    var simulateRequestURI = "http://localhost/no-way";
//    var requestUri = new URI(simulateRequestURI);
//    HttpHeaders headers = mock(HttpHeaders.class);
//    when(request.getURI()).thenReturn(new URI(simulateRequestURI));
//    when(request.getMethod()).thenReturn(HttpMethod.POST);
//
//    when(exchange.getAttributeOrDefault(anyString(), anySet()))
//        .thenReturn(Set.of(new URI("login")));
//    when(exchange.getAttribute(GATEWAY_ROUTE_ATTR)).thenReturn(route);
//    when(exchange.getAttribute(GATEWAY_REQUEST_URL_ATTR)).thenReturn(requestUri);
//    when(request.getHeaders()).thenReturn(headers);
//    when(headers.get(anyString())).thenReturn(null);
//
//    GatewayFilter filter = config.apply(new AuthFilterConfig.Config());
//    Assertions.assertThrows(NullPointerException.class, () -> filter.filter(exchange, chain));
//  }
//}
