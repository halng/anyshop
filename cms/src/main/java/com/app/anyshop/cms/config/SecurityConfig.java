
package com.app.anyshop.cms.config;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configuration.WebSecurityCustomizer;
import org.springframework.security.config.annotation.web.configurers.AbstractHttpConfigurer;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.web.SecurityFilterChain;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;
import org.springframework.web.servlet.config.annotation.CorsRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;

@Configuration
public class SecurityConfig {
  private final RequestFilter requestFilter;

  @Autowired
  SecurityConfig(RequestFilter requestFilter) {
    this.requestFilter = requestFilter;
  }
  @Bean
  public WebSecurityCustomizer webSecurityCustomizer() {
    return webSecurity ->
        webSecurity
            .ignoring()
            .requestMatchers(
                "/actuator/prometheus",
                "/actuator/prometheus/**",
                "/swagger-ui",
                "/swagger-ui/**",
                "/error",
                "/v3/api-docs/**");
  }

  @Bean
  public WebMvcConfigurer corsConfigurer() {
    return new WebMvcConfigurer() {
      @Override
      public void addCorsMappings(CorsRegistry corsRegistry) {
        corsRegistry.addMapping("/**").allowedMethods("*").allowedOrigins("*").allowedHeaders("*");
      }
    };
  }

  @Bean
  public SecurityFilterChain filterChain(HttpSecurity http) throws Exception {
    http.sessionManagement(sess -> sess.sessionCreationPolicy(SessionCreationPolicy.STATELESS))
        .authorizeHttpRequests(
            req ->
                req.requestMatchers("/swagger-ui.html", "/swagger-ui/**", "/api-docs/**", "/error")
                    .permitAll()
                    .anyRequest()
                    .authenticated());

    http.csrf(AbstractHttpConfigurer::disable).cors(AbstractHttpConfigurer::disable);
    http.addFilterBefore(requestFilter, UsernamePasswordAuthenticationFilter.class);
    return http.build();
  }
}
