package org.wid.userservice.dto;

import org.springframework.http.HttpStatus;

public record ErrorResponseDto(HttpStatus status, String message) {
}
