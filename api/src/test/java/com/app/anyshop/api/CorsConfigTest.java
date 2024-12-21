//package com.app.anyshop.api;
//
//import static org.junit.jupiter.api.Assertions.*;
//
//import com.app.anyshop.api.config.CorsConfig;
//import java.util.ArrayList;
//import java.util.Arrays;
//import org.junit.jupiter.api.BeforeEach;
//import org.junit.jupiter.api.Test;
//import org.springframework.http.HttpHeaders;
//import org.springframework.http.HttpMethod;
//import org.springframework.test.web.reactive.server.EntityExchangeResult;
//import org.springframework.test.web.reactive.server.WebTestClient;
//
//class CorsConfigTest {
//  private final CorsConfig corsConfig = new CorsConfig();
//  private WebTestClient webClient;
//
//  @BeforeEach
//  void setUp() {
//    webClient =
//        WebTestClient.bindToController(new MockController())
//            .webFilter(corsConfig.corsWebFilter())
//            .build();
//  }
//
//  @Test
//  void testVerifyCorsWebFilter_nonCorsRequest() {
//    EntityExchangeResult<?> result =
//        webClient
//            .get()
//            .uri("/mock/hello")
//            .exchange()
//            .expectStatus()
//            .isOk()
//            .expectBody()
//            .returnResult();
//    var headers = result.getResponseHeaders();
//    assertNotNull(headers);
//
//    assertNull(headers.getFirst(HttpHeaders.ACCESS_CONTROL_ALLOW_ORIGIN));
//    assertTrue(headers.getAccessControlAllowMethods().isEmpty());
//    assertTrue(headers.getAccessControlAllowHeaders().isEmpty());
//    assertEquals(-1, headers.getAccessControlMaxAge());
//  }
//
//  @Test
//  void testVerifyCorsWebFilter_corsRequest() {
//    EntityExchangeResult<?> result =
//        webClient
//            .get()
//            .uri("http://localhost:8080/mock/hello")
//            .header(HttpHeaders.ORIGIN, "http://localhost:8081")
//            .exchange()
//            .expectStatus()
//            .isOk()
//            .expectBody()
//            .returnResult();
//    var headers = result.getResponseHeaders();
//    assertNotNull(headers);
//
//    assertEquals("*", headers.getFirst(HttpHeaders.ACCESS_CONTROL_ALLOW_ORIGIN));
//    assertEquals(
//        new ArrayList<>(
//            Arrays.asList(
//                HttpMethod.GET,
//                HttpMethod.POST,
//                HttpMethod.PUT,
//                HttpMethod.DELETE,
//                HttpMethod.OPTIONS)),
//        headers.getAccessControlAllowMethods());
//    assertEquals(
//        new ArrayList<>(
//            Arrays.asList(
//                "authorization",
//                "Content-Type",
//                "Authorization",
//                "credential",
//                "X-API-SECRET-TOKEN",
//                "X-API-USER-ID")),
//        headers.getAccessControlAllowHeaders());
//    assertEquals(3600, headers.getAccessControlMaxAge());
//  }
//
//  @Test
//  void testVerifyCorsWebFilter_corsRequestWithOption() {
//    EntityExchangeResult<?> result =
//        webClient
//            .options()
//            .uri("http://localhost:8080/mock/hello")
//            .header(HttpHeaders.ORIGIN, "http://localhost:8081")
//            .exchange()
//            .expectStatus()
//            .isOk()
//            .expectBody()
//            .returnResult();
//  }
//}
