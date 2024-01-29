package org.wid.userservice.dto.oauth2.resource;

public record GoogleUserDto(
    String id,
    String email,
    Boolean verifiedEmail,
    String name,
    String familyName,
    String givenName,
    String picture,
    String locale) {
}
