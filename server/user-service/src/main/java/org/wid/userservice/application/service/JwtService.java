package org.wid.userservice.application.service;

import java.time.Instant;
import java.util.Date;

import org.springframework.stereotype.Service;
import org.wid.userservice.application.dto.auth.JwtDto;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.port.driving.JwtServicePort;

import io.jsonwebtoken.Jwts;

@Service
public class JwtService implements JwtServicePort {

  @Override
  public JwtDto generateTokens(UserDto userDto) {
    return new JwtDto(
        generateAccessToken(userDto),
        generateRefreshToken(userDto));
  }

  @Override
  public JwtDto refresh(String refreshToken) {
    // TODO Auto-generated method stub
    throw new UnsupportedOperationException("Unimplemented method 'refresh'");
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
