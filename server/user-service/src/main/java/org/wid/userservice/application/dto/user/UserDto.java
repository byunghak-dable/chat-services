package org.wid.userservice.application.dto.user;

import com.fasterxml.jackson.databind.PropertyNamingStrategies.SnakeCaseStrategy;
import com.fasterxml.jackson.databind.annotation.JsonNaming;
import org.wid.userservice.domain.entity.User.LoginType;

@JsonNaming(SnakeCaseStrategy.class)
public record UserDto(String id, String email, String name, String profile, LoginType loginType) {}
