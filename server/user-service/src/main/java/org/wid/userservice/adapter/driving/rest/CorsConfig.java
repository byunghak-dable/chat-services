package org.wid.userservice.adapter.driving.rest;

import java.util.List;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.web.cors.CorsConfiguration;
import org.springframework.web.cors.reactive.CorsWebFilter;
import org.springframework.web.cors.reactive.UrlBasedCorsConfigurationSource;

@Configuration
public class CorsConfig {

  @Bean
  public CorsWebFilter corsWebFilter() {
    CorsConfiguration corsConfig = new CorsConfiguration();
    UrlBasedCorsConfigurationSource source = new UrlBasedCorsConfigurationSource();

    corsConfig.setAllowedOrigins(List.of("*"));
    corsConfig.setAllowedMethods(List.of("GET", "POST", "PUT", "DELETE"));
    corsConfig.setAllowedHeaders(List.of("Content-Type", "Authorization"));
    // corsConfig.setAllowCredentials(true);

    source.registerCorsConfiguration("/**", corsConfig);

    return new CorsWebFilter(source);
  }
}
