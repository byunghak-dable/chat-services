package org.wid.userservice.adapter.driven.client.oauth2.dto;

import com.fasterxml.jackson.databind.PropertyNamingStrategies.SnakeCaseStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import java.util.Optional;

@JsonNaming(SnakeCaseStrategy.class)
public record Oauth2TokenDto(
    String tokenType, String scope, String accessToken, Optional<String> refreshToken) {}
