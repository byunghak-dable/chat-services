package org.wid.userservice.adapter.driven.oauth2.dto.github;

public record GithubTokenRequestDto(
    String clientId, String clientSecret, String redirectUri, String code) {}
