package org.wid.userservice.service;

import java.time.Instant;
import java.util.Date;

import org.springframework.stereotype.Service;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.port.primary.JwtServicePort;

import io.jsonwebtoken.Jwts;

@Service
public class JwtService implements JwtServicePort {

  @Override
  public String createAccessToken(UserDto userDto) {
    long expiredSeconds = 5 * 60;

    return createToken(userDto, expiredSeconds);
  }

  @Override
  public String createRefreshToken(UserDto userDto) {
    long expiredSecondsToAdd = 14 * 24 * 60 * 60;

    return createToken(userDto, expiredSecondsToAdd);
  }

  private String createToken(UserDto userDto, long expirationDurationSeconds) {
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
