package org.wid.userservice.adapter.driven.oauth2.dto.google;

public record GoogleUserDto(
    String id,
    String email,
    Boolean verifiedEmail,
    String name,
    String familyName,
    String givenName,
    String picture,
    String locale) {}
