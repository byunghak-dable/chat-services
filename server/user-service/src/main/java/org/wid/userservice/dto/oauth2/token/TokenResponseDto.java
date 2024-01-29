package org.wid.userservice.dto.oauth2.token;

import java.util.Optional;

import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(PropertyNamingStrategies.SnakeCaseStrategy.class)
public record TokenResponseDto(
    String tokenType,
    String accessToken,
    String scope,
    Optional<String> refreshToken) {
}
