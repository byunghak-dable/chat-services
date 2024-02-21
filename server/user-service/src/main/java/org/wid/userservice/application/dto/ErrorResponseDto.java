package org.wid.userservice.application.dto;

import org.springframework.http.HttpStatus;

public record ErrorResponseDto(HttpStatus status, String message) {
}
