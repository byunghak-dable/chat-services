package org.wid.userservice.adapter.primary.rest.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.dto.user.UserDto;
import org.wid.userservice.port.primary.UserServicePort;

import lombok.RequiredArgsConstructor;
import reactor.core.publisher.Mono;

@RestController
@RequestMapping("/api/v1")
@RequiredArgsConstructor
public class UserController {
  private final UserServicePort userService;

  @GetMapping("/user/{userId}")
  public Mono<UserDto> getUser(@PathVariable String userId) {
    return userService.getUser(userId);
  }
}
