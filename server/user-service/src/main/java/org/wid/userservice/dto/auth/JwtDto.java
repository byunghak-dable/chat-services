package org.wid.userservice.dto.auth;

import com.fasterxml.jackson.databind.PropertyNamingStrategies.SnakeCaseStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(SnakeCaseStrategy.class)
public record JwtDto(
    String accessToken,
    String refreshToken) {
}
