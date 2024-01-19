package org.wid.userservice.dto.user;

import jakarta.validation.constraints.Email;
import jakarta.validation.constraints.NotNull;

public record RegisterUserDto(
    @Email @NotNull String email,
    @NotNull String password,
    @NotNull String name,
    String gender,
    String thumbnailUrl) {
}
