package com.app.anyshop.cms.annotation;

import com.app.anyshop.cms.annotation.validator.ValidActionValidator;
import jakarta.validation.Constraint;
import jakarta.validation.Payload;
import java.lang.annotation.ElementType;
import java.lang.annotation.Retention;
import java.lang.annotation.RetentionPolicy;
import java.lang.annotation.Target;

@Constraint(validatedBy = ValidActionValidator.class)
@Target({ElementType.PARAMETER, ElementType.FIELD})
@Retention(RetentionPolicy.RUNTIME)
public @interface ValidAction {
  String message() default "Invalid action.";

  Class<?>[] groups() default {};

  Class<? extends Payload>[] payload() default {};
}
