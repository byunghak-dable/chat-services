package org.wid.userservice.adapter.driven.config;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;

import lombok.Data;

@Configuration
public class Oauth2ClientConfig {

  @Bean
  @ConfigurationProperties("oauth2.provider.github")
  public OAuth2ClientProperties githubProperties() {
    return new OAuth2ClientProperties();
  }

  @Bean
  @ConfigurationProperties("oauth2.provider.google")
  public OAuth2ClientProperties googleProperties() {
    return new OAuth2ClientProperties();
  }

  @Data
  public static class OAuth2ClientProperties {
    private String clientId;
    private String clientSecret;
    private String redirectUri;
    private String tokenUri;
    private String resourceUri;
  }
}
