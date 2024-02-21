package org.wid.userservice.application.dto.oauth2.token;

public record GithubTokenRequestDto(
    String clientId,
    String clientSecret,
    String redirectUri,
    String code) {
}
