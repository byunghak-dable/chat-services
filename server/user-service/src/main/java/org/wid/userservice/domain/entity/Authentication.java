package org.wid.userservice.domain.entity;

import io.jsonwebtoken.Jwts;
import java.time.Instant;
import java.util.Date;
import lombok.Getter;

@Getter
public class Authentication {
  public final String subject;
  public final Date issuedAt;
  public final Date expiration;

  private Authentication(String subject, long expirationSeconds) {
    Instant currentUtc = Instant.now();
    Instant expirationInstant = currentUtc.plusSeconds(expirationSeconds);

    this.subject = subject;
    this.issuedAt = Date.from(currentUtc);
    this.expiration = Date.from(expirationInstant);
  }

  // TODO: require implementation
  public static Authentication from(String refreshToken) {
    return new Authentication("", 5 * 60);
  }

  public static Authentication accessToken(String subject) {
    return new Authentication(subject, 5 * 60);
  }

  public static Authentication refreshToken(String subject) {
    return new Authentication(subject, 14 * 24 * 60 * 60);
  }

  public String toJsonWebToken() {
    return Jwts.builder()
        .issuer("me")
        .subject(this.subject)
        .issuedAt(this.issuedAt)
        .expiration(this.expiration)
        .compact();
  }
}
