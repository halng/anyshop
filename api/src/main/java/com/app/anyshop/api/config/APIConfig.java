package com.app.anyshop.api.config;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.cloud.gateway.filter.factory.DedupeResponseHeaderGatewayFilterFactory;
import org.springframework.cloud.gateway.route.RouteLocator;
import org.springframework.cloud.gateway.route.builder.RouteLocatorBuilder;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class APIConfig {
  private static final String APP_IAM_ID = "iam";

  @Value("${host.iam}")
  private String IAM_HOST;

  private static final String IAM_PATH = "/api/v1/iam/**";
  private static final String IAM_PATH_REGEX = "/api/v1/iam(?<segment>/?.*)";
  private static final String IAM_PATH_REPLACEMENT = "/api/v1/iam${segment}";

  // others host

  @Bean
  public RouteLocator routes(
      RouteLocatorBuilder builder,
      AuthFilterConfig authFilterConfig,
      DedupeResponseHeaderGatewayFilterFactory dedupe) {
    return builder
        .routes()
        .route(
            APP_IAM_ID,
            predicateSpec ->
                predicateSpec
                    .path(IAM_PATH)
                    .filters(
                        filter ->
                            filter
                                .rewritePath(IAM_PATH_REGEX, IAM_PATH_REPLACEMENT)
                                .filter(dedupe.apply(dedupeResponseHeader()))
                                .filter(authFilterConfig.apply(new AuthFilterConfig.Config())))
                    .uri(IAM_HOST))
        .build();
  }

  @Bean
  public DedupeResponseHeaderGatewayFilterFactory.Config dedupeResponseHeader() {
    DedupeResponseHeaderGatewayFilterFactory.Config ret =
        new DedupeResponseHeaderGatewayFilterFactory.Config();
    ret.setStrategy(DedupeResponseHeaderGatewayFilterFactory.Strategy.RETAIN_FIRST);
    ret.setName("Access-Control-Allow-Origin");
    return ret;
  }

}
