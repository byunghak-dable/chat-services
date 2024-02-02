package org.wid.userservice.service;

import java.util.Date;

import org.springframework.stereotype.Service;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.port.primary.JwtServicePort;

import io.jsonwebtoken.Jwts;

@Service
public class JwtService implements JwtServicePort {

  @Override
  public String createAccessToken(UserDto userDto) {
    return createJwt(userDto.id(), null);
  }

  @Override
  public String createRefreshToken(UserDto userDto) {
    return createJwt(userDto.id(), null);
  }

  private String createJwt(String subject, Date expiredDate) {
    return Jwts.builder()
        .issuer("me")
        .subject(subject)
        .issuedAt(new Date())
        .expiration(expiredDate)
        .compact();
  }
}
