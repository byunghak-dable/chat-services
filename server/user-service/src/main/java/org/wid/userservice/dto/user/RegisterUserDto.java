package org.wid.userservice.dto.user;

import jakarta.validation.constraints.Email;

public record RegisterUserDto(
    @Email String email,
    String password,
    String name,
    String gender,
    String thumbnailUrl) {
}
