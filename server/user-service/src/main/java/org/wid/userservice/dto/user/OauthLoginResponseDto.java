package org.wid.userservice.dto.user;

public record OauthLoginResponseDto(String accessToken, String refreshToken) {
}
