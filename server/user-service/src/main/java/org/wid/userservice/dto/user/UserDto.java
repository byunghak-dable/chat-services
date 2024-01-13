package org.wid.userservice.dto.user;

public record UserDto(
    Long id,
    String email,
    String name,
    Boolean gender,
    String thumbnailUrl) {
}
