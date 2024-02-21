package org.wid.userservice.application.dto.oauth2.token;

public record GoogleTokenRequestDto(
    String grantType,
    String clientId,
    String clientSecret,
    String code,
    String redirectUri) {

  public GoogleTokenRequestDto(String clientId, String clientSecret, String redirectUri, String code) {
    this("authorization_code", clientId, clientSecret, code, redirectUri);
  }
}
