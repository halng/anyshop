//package com.app.anyshop.api;
//
//import static org.junit.jupiter.api.Assertions.*;
//import static org.mockito.Mockito.*;
//
//import com.app.anyshop.api.config.APIConfig;
//import org.junit.jupiter.api.BeforeEach;
//import org.junit.jupiter.api.Test;
//import org.mockito.Mock;
//import org.mockito.MockitoAnnotations;
//import org.springframework.cloud.gateway.filter.factory.DedupeResponseHeaderGatewayFilterFactory;
//import org.springframework.context.ConfigurableApplicationContext;
//import org.springframework.test.util.ReflectionTestUtils;
//import org.springframework.web.reactive.function.client.WebClient;
//
//class APIConfigTest {
//
//  private APIConfig apiConfig;
//
//  //  @Mock private RouteLocatorBuilder routeLocatorBuilder;
//  //
////  private RouteLocatorBuilder.Builder routesBuilder;
//  //  @Mock private Route.Builder routeBuilder;
//  //
////  private AuthFilterConfig authFilterConfig;
//
////  @Mock private WebClient.Builder webClientBuilder;
//
//  //  @Mock private DedupeResponseHeaderGatewayFilterFactory
//  // dedupeResponseHeaderGatewayFilterFactory;
//
//  @Mock private ConfigurableApplicationContext context;
//
//  //  @Mock private GatewayFilter gatewayFilter;
//
//  @BeforeEach
//  void setup() {
//    MockitoAnnotations.openMocks(this);
//    apiConfig = new APIConfig();
//    //    authFilterConfig = new AuthFilterConfig(webClientBuilder);
//    //    routesBuilder = new RouteLocatorBuilder.Builder(context);
//    ReflectionTestUtils.setField(apiConfig, "IAM_HOST", "http://iam.example.com");
//    ReflectionTestUtils.setField(apiConfig, "CMS_HOST", "http://cms.example.com");
//
//    //    apiConfig.IAM_HOST = "http://iam.example.com";
//    //    apiConfig.CMS_HOST = "http://cms.example.com";
//  }
//
//  @Test
//  void testDedupeResponseHeader() {
//    DedupeResponseHeaderGatewayFilterFactory.Config config = apiConfig.dedupeResponseHeader();
//
//    assertNotNull(config);
//    assertEquals(
//        DedupeResponseHeaderGatewayFilterFactory.Strategy.RETAIN_FIRST, config.getStrategy());
//    assertEquals("Access-Control-Allow-Origin", config.getName());
//  }
//  //
//  //  @Test
//  //  void testRoutesConfiguration() {
//  //    // Mock the routes() method on RouteLocatorBuilder
//  //    when(routeLocatorBuilder.routes()).thenReturn(routesBuilder);
//  //    when(routesBuilder.route(
//  //            anyString(),
//  //            argThat(fn -> {
//  //              // Simulate the behavior of the function
//  //              Route.Builder routeBuilderMock = mock(Route.Builder.class);
//  //              when(routeBuilderMock.filters(gatewayFilter)).thenReturn(routeBuilderMock);
//  //              when(routeBuilderMock.uri(anyString())).thenReturn(routeBuilderMock);
//  //              fn.apply()
//  //              return true;
//  //            })))
//  //            .thenReturn(routesBuilder);
//  //    DedupeResponseHeaderGatewayFilterFactory.Config config = apiConfig.dedupeResponseHeader();
//  //    when(dedupeResponseHeaderGatewayFilterFactory.apply(config)).thenReturn(gatewayFilter);
//  //
//  //    // Call the method under test
//  //    RouteLocator routeLocator =
//  //        apiConfig.routes(
//  //            routeLocatorBuilder, authFilterConfig, dedupeResponseHeaderGatewayFilterFactory);
//  //
//  //    // Verify that routes() was called
//  //    verify(routeLocatorBuilder, times(1)).routes();
//  //    assertNotNull(routeLocator);
//  //  }
//  //
//  //  @Test
//  //  void testIAMRouteConfiguration() {
//  //    when(routeLocatorBuilder.routes()).thenReturn(routesBuilder);
//  //
//  //    // Capture the lambda passed to `route` for IAM route
//  //    ArgumentCaptor<Function<PredicateSpec, Buildable<Route>>> captor =
//  //        ArgumentCaptor.forClass(Function.class);
//  //
//  //    when(routesBuilder.route(eq("iam"), captor.capture())).thenReturn(routesBuilder);
//  //
//  //    // Call the method under test
//  //    apiConfig.routes(
//  //        routeLocatorBuilder, authFilterConfig, dedupeResponseHeaderGatewayFilterFactory);
//  //
//  //    // Verify that the IAM route was defined
//  //    verify(routesBuilder, times(1)).route(eq("iam"), any());
//  //
//  //    // Validate the captured lambda's behavior
//  //    PredicateSpec predicateSpec = mock(PredicateSpec.class);
//  //    BooleanSpec booleanSpec = mock(BooleanSpec.class);
//  //    when(predicateSpec.path("/api/v1/iam/**")).thenReturn(booleanSpec);
//  //
//  //    captor.getValue().apply(predicateSpec);
//  //
//  //    verify(predicateSpec, times(1)).path("/api/v1/iam/**");
//  //    verify(booleanSpec, times(1)).filters(any());
//  //  }
//  //
//  //  @Test
//  //  void testCMSRouteConfiguration() {
//  //    when(routeLocatorBuilder.routes()).thenReturn(routesBuilder);
//  //
//  //    // Capture the lambda passed to `route` for CMS route
//  //    ArgumentCaptor<Function<PredicateSpec, Buildable<Route>>> captor =
//  //        ArgumentCaptor.forClass(Function.class);
//  //
//  //    when(routesBuilder.route(eq("cms"), captor.capture())).thenReturn(routesBuilder);
//  //
//  //    // Call the method under test
//  //    apiConfig.routes(
//  //        routeLocatorBuilder, authFilterConfig, dedupeResponseHeaderGatewayFilterFactory);
//  //
//  //    // Verify that the CMS route was defined
//  //    verify(routesBuilder, times(1)).route(eq("cms"), any());
//  //
//  //    // Validate the captured lambda's behavior
//  //    PredicateSpec predicateSpec = mock(PredicateSpec.class);
//  //    BooleanSpec routeBuilder = mock(BooleanSpec.class);
//  //    when(predicateSpec.path("/api/v1/cms/**")).thenReturn(routeBuilder);
//  //
//  //    captor.getValue().apply(predicateSpec);
//  //
//  //    verify(predicateSpec, times(1)).path("/api/v1/cms/**");
//  //    verify(routeBuilder, times(1)).filters(any());
//  //  }
//}
