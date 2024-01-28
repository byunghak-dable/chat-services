package org.wid.userservice.adapter.primary.rest;

import java.util.List;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.annotation.web.reactive.EnableWebFluxSecurity;
import org.springframework.security.config.web.server.ServerHttpSecurity;
import org.springframework.security.web.server.SecurityWebFilterChain;
import org.springframework.web.cors.CorsConfiguration;
import org.springframework.web.cors.reactive.CorsConfigurationSource;
import org.springframework.web.cors.reactive.UrlBasedCorsConfigurationSource;
import org.wid.userservice.port.primary.AuthServicePort;

import lombok.RequiredArgsConstructor;

@Configuration
@EnableWebFluxSecurity
@RequiredArgsConstructor
public class SecurityConfig {

  private final AuthServicePort authService;

  @Bean
  public SecurityWebFilterChain basicFilterChain(ServerHttpSecurity http) throws Exception {
    return http
        .csrf(csrf -> csrf.disable())
        .authorizeExchange(exchanges -> exchanges
            .pathMatchers("/auth/**").permitAll()
            // .pathMatchers("/api/**").authenticated())
            // TODO: for test
            .pathMatchers("/api/**").permitAll())
        .build();
  }

  @Bean
  CorsConfigurationSource corsConfigurationSource() {
    CorsConfiguration configuration = new CorsConfiguration();
    UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();

    configuration.setAllowedOriginPatterns(List.of("*://localhost:*"));
    configuration.setAllowedMethods(List.of("GET", "POST", "PUT", "DELETE"));
    configuration.setAllowedHeaders(List.of("Content-Type", "Authorization"));
    configuration.setAllowCredentials(true);
    source.registerCorsConfiguration("/**", configuration);

    return source;
  }
}
