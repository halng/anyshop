//package com.app.anyshop.api.config;
//
//import org.springframework.context.annotation.Bean;
//import org.springframework.context.annotation.Configuration;
//import org.springframework.http.HttpHeaders;
//import org.springframework.http.HttpMethod;
//import org.springframework.http.HttpStatus;
//import org.springframework.http.server.reactive.ServerHttpRequest;
//import org.springframework.http.server.reactive.ServerHttpResponse;
//import org.springframework.web.cors.reactive.CorsUtils;
//import org.springframework.web.server.ServerWebExchange;
//import org.springframework.web.server.WebFilter;
//import org.springframework.web.server.WebFilterChain;
//import reactor.core.publisher.Mono;
//
//@Configuration
//public class CorsConfig {
//    private static final String ALLOWED_HEADERS = "authorization, Content-Type, Authorization, credential, X-API-SECRET-TOKEN, X-API-USER-ID";
//    private static final String ALLOWED_METHODS = "GET, POST, PUT, DELETE, OPTIONS";
//    private static final String ALLOWED_ORIGIN = "*";
//    private static final String MAX_AGE = "3600";
//
//    @Bean
//    public WebFilter corsWebFilter() {
//        return (ServerWebExchange sse, WebFilterChain wfc) -> {
//
//            ServerHttpRequest request = sse.getRequest();
//            if (CorsUtils.isCorsRequest(request)) {
//                ServerHttpResponse response = sse.getResponse();
//                HttpHeaders headers = response.getHeaders();
//                headers.set(HttpHeaders.ACCESS_CONTROL_ALLOW_ORIGIN, ALLOWED_ORIGIN);
//                headers.set(HttpHeaders.ACCESS_CONTROL_ALLOW_METHODS, ALLOWED_METHODS);
//                headers.set(HttpHeaders.ACCESS_CONTROL_ALLOW_HEADERS, ALLOWED_HEADERS);
//                headers.set(HttpHeaders.ACCESS_CONTROL_MAX_AGE, MAX_AGE);
//                if (request.getMethod() == HttpMethod.OPTIONS) {
//                    response.setStatusCode(HttpStatus.OK);
//                    return Mono.empty();
//                }
//            }
//            return wfc.filter(sse);
//        };
//    }
//}
