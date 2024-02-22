package org.wid.userservice.application.service;

import java.time.Instant;
import java.util.Date;

import org.springframework.stereotype.Service;
import org.wid.userservice.application.dto.auth.AccessTokenDto;
import org.wid.userservice.application.dto.auth.AuthenticationTokensDto;
import org.wid.userservice.application.dto.user.UserDto;

import io.jsonwebtoken.Jwts;

interface TokenService {
  AuthenticationTokensDto generateTokens(UserDto userDto);

  AccessTokenDto generateAccessToken(String refreshToken);
}

@Service
public class JwtService implements TokenService {

  @Override
  public AuthenticationTokensDto generateTokens(UserDto userDto) {
    return new AuthenticationTokensDto(
        generateAccessToken(userDto),
        generateRefreshToken(userDto));
  }

  @Override
  public AccessTokenDto generateAccessToken(String refreshToken) {
    // TODO Auto-generated method stub
    throw new UnsupportedOperationException("Unimplemented method 'generateAccessToken'");
  }

  private String generateAccessToken(UserDto userDto) {
    long expiredSeconds = 5 * 60;

    return generateToken(userDto, expiredSeconds);
  }

  private String generateRefreshToken(UserDto userDto) {
    long expiredSecondsToAdd = 14 * 24 * 60 * 60;

    return generateToken(userDto, expiredSecondsToAdd);
  }

  private String generateToken(UserDto userDto, long expirationDurationSeconds) {
    Instant currentUtc = Instant.now();
    Instant expirationInstant = currentUtc.plusSeconds(expirationDurationSeconds);

    return Jwts.builder()
        .issuer("me")
        .subject(userDto.id())
        .issuedAt(Date.from(currentUtc))
        .expiration(Date.from(expirationInstant))
        .compact();
  }
}
