package org.wid.userservice.dto.oauth2;

import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(PropertyNamingStrategies.SnakeCaseStrategy.class)
public record TokenRequestDto(String clientId, String clientSecret, String code, String grantType,
    String redirectUri) {

  public TokenRequestDto(String clientId, String clientSecret, String code) {
    this(clientId, clientSecret, code, "authorization_code", "http://localhost:3000");
  }
}
