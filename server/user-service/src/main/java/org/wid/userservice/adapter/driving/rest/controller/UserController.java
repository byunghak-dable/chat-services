package org.wid.userservice.adapter.driving.rest.controller;

import java.util.Map;
import lombok.RequiredArgsConstructor;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.application.dto.user.UserDto;
import org.wid.userservice.domain.entity.User.LoginType;
import org.wid.userservice.port.driving.UserServicePort;
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

  @PostMapping("/user")
  public Mono<UserDto> getTestUser(@RequestBody Map<String, String> t) {
    return Mono.just(new UserDto("1", "email", "firstName", "thumnail", LoginType.GOOGLE));
  }
}
