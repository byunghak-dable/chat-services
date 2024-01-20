package org.wid.userservice.dto.user;

import com.fasterxml.jackson.databind.PropertyNamingStrategies;
import com.fasterxml.jackson.databind.annotation.JsonNaming;

@JsonNaming(PropertyNamingStrategies.SnakeCaseStrategy.class)
public record UserDto(
    Long id,
    String email,
    String firstName,
    String lastName,
    Boolean gender,
    String thumbnailUrl) {
}
