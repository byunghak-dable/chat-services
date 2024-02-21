package org.wid.userservice.application.dto.user;

import org.wid.userservice.domain.entity.User.LoginType;

import com.fasterxml.jackson.databind.PropertyNamingStrategies.SnakeCaseStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(SnakeCaseStrategy.class)
public record UserDto(
    String id,
    String email,
    String name,
    String profile,
    LoginType loginType) {
}
