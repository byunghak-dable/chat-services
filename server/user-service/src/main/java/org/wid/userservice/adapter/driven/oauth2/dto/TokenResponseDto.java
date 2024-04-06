package org.wid.userservice.adapter.driven.oauth2.dto;

import java.util.Optional;

import com.fasterxml.jackson.databind.PropertyNamingStrategies.SnakeCaseStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(SnakeCaseStrategy.class)
public record TokenResponseDto(
    String tokenType, String scope, String accessToken, Optional<String> refreshToken) {}
