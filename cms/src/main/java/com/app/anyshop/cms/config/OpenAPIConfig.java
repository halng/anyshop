package com.app.anyshop.cms.config;

import io.swagger.v3.oas.models.Components;
import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.info.Info;
import io.swagger.v3.oas.models.security.SecurityRequirement;
import io.swagger.v3.oas.models.security.SecurityScheme;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Value;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

@Configuration
public class OpenAPIConfig {

  private final String moduleName;

  private final String apiVersion;

  public OpenAPIConfig(
      @Value("${module-name:Module}") String moduleName,
      @Value("${api-version:v1}") String apiVersion) {
    this.moduleName = moduleName;
    this.apiVersion = apiVersion;
  }

  @Bean
  public OpenAPI customAPI() {
    final String userIdSchema = "X-API-USER-ID: ";
    final String userSchema = "X-API-USER: ";
    final String roleSchema = "X-API-USER-ROLE: ";
    final String apiTitle = String.format("%s API", StringUtils.capitalize(this.moduleName));

    return new OpenAPI()
        .addSecurityItem(new SecurityRequirement().addList(userSchema))
        .components(
            new Components()
                .addSecuritySchemes(
                    userSchema,
                    new SecurityScheme()
                        .name(userSchema)
                        .type(SecurityScheme.Type.HTTP)
                        .scheme("X-API-USER")))
        .addSecurityItem(new SecurityRequirement().addList(roleSchema))
        .components(
            new Components()
                .addSecuritySchemes(
                    roleSchema,
                    new SecurityScheme()
                        .name(roleSchema)
                        .type(SecurityScheme.Type.HTTP)
                        .scheme("X-API-USER-ROLE")))
        .addSecurityItem(new SecurityRequirement().addList(userIdSchema))
        .components(
            new Components()
                .addSecuritySchemes(
                    userIdSchema,
                    new SecurityScheme()
                        .name(userIdSchema)
                        .type(SecurityScheme.Type.HTTP)
                        .scheme("X-API-USER-ID")))
        .info(new Info().title(apiTitle).version(apiVersion));
  }
}
