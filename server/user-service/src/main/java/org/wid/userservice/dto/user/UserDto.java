package org.wid.userservice.dto.user;

import org.wid.userservice.entity.entity.User.LoginType;

import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(PropertyNamingStrategies.SnakeCaseStrategy.class)
public record UserDto(
    String id,
    String email,
    String name,
    String profile,
    LoginType loginType) {
}
