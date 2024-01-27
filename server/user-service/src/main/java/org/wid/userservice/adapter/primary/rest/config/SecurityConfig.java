package org.wid.userservice.adapter.primary.rest.config;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.security.config.Customizer;
import org.springframework.security.config.annotation.web.reactive.EnableWebFluxSecurity;
import org.springframework.security.config.web.server.ServerHttpSecurity;
import org.springframework.security.web.server.SecurityWebFilterChain;
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
        .cors(Customizer.withDefaults())
        .authorizeExchange(exchanges -> exchanges
            .pathMatchers("/auth/**").permitAll()
            .pathMatchers("/api/**").authenticated())
        .csrf(csrf -> csrf.disable())
        .build();
  }
}
