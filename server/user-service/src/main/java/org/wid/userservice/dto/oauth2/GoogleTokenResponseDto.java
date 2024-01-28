package org.wid.userservice.dto.oauth2;

import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(PropertyNamingStrategies.SnakeCaseStrategy.class)
public record GoogleTokenResponseDto(
    String tokenType,
    String accessToken,
    String refreshToken,
    String scope,
    int expiresIn) {
}
