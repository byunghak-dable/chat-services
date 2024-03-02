package org.wid.userservice.application.service;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;
import org.springframework.boot.test.autoconfigure.web.servlet.WebMvcTest;
import org.springframework.test.web.servlet.MockMvc;
import org.wid.userservice.adapter.driving.rest.controller.UserController;

@WebMvcTest(UserController.class)
public class AuthFacadeServiceTest {

  private MockMvc mockMvc;

  @BeforeEach
  void beforeEach() {
    System.out.println("testing");
  }

  @DisplayName("testing method")
  @Test
  public void testing() {

  }
}
