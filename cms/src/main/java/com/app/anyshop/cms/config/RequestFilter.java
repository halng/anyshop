package com.app.anyshop.cms.config;

import jakarta.servlet.FilterChain;
import jakarta.servlet.ServletException;
import jakarta.servlet.http.HttpServletRequest;
import jakarta.servlet.http.HttpServletResponse;
import lombok.RequiredArgsConstructor;
import org.apache.commons.lang3.StringUtils;
import org.apache.commons.lang3.tuple.ImmutablePair;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Component;
import org.springframework.web.filter.OncePerRequestFilter;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

@Component
@RequiredArgsConstructor
public class RequestFilter extends OncePerRequestFilter {
  private static final Logger logger = LoggerFactory.getLogger(RequestFilter.class);
  private static final String USER_ID_KEY = "X-API-USER-ID";
  private static final String USER_NAME_KEY = "X-API-USER";
  private static final String USER_ROLE_KEY = "X-API-USER-ROLE";

  @Override
  protected void doFilterInternal(
      HttpServletRequest request, HttpServletResponse response, FilterChain filterChain)
      throws ServletException, IOException {
    logger.info("RequestFilter.doFilterInternal: Start verifying incoming request");
    String username = request.getHeader(USER_NAME_KEY);
    String role = request.getHeader(USER_ROLE_KEY);
    String userId = request.getHeader(USER_ID_KEY);

    UsernamePasswordAuthenticationToken credentials = null;

    if (StringUtils.isNotEmpty(username)
        && StringUtils.isNotEmpty(role)
        && StringUtils.isNotEmpty(userId)) {
      List<GrantedAuthority> authorities = new ArrayList<>();
      authorities.add(new SimpleGrantedAuthority(role));
      ImmutablePair<String, String> principles = new ImmutablePair<>(username, userId);
      credentials = new UsernamePasswordAuthenticationToken(principles, null, authorities);

      logger.info("Authentication successful");
    } else {
      logger.error("Can not validate user!");
    }

    SecurityContextHolder.getContext().setAuthentication(credentials);
    filterChain.doFilter(request, response);
  }
}
