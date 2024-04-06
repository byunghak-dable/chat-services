package org.wid.userservice.domain.entity;

import io.jsonwebtoken.Jwts;
import java.time.Instant;
import java.util.Date;
import lombok.Getter;

@Getter
public class Authentication {
  public final String userId;
  public final Date issuedAt;
  public final Date expiration;

  private Authentication(String userId, long expirationSeconds) {
    Instant currentUtc = Instant.now();
    Instant expirationInstant = currentUtc.plusSeconds(expirationSeconds);

    this.userId = userId;
    this.issuedAt = Date.from(currentUtc);
    this.expiration = Date.from(expirationInstant);
  }

  // TODO: require implementation
  public static Authentication renewAccess(String refreshToken) {
    return new Authentication("", 5 * 60);
  }

  public static Authentication createAccess(String userId) {
    return new Authentication(userId, 5 * 60);
  }

  public static Authentication createRefresh(String userId) {
    return new Authentication(userId, 14 * 24 * 60 * 60);
  }

  public String toJsonWebToken() {
    return Jwts.builder()
        .issuer("chat")
        .subject(this.userId)
        .issuedAt(this.issuedAt)
        .expiration(this.expiration)
        .compact();
  }
}
