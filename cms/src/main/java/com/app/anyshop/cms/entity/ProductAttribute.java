/*
 * *****************************************************************************************
 * Copyright 2024 By Hal Nguyen
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * *****************************************************************************************
 */

package com.app.anyshop.cms.entity;

import jakarta.persistence.*;
import java.util.List;
import lombok.*;

@Getter
@Setter
@Table(name = "product_attributes")
@Entity
@Builder
@AllArgsConstructor
@NoArgsConstructor
public class ProductAttribute extends Audit {
  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  private String id;

  private String name;
  private String description;

  @Enumerated(EnumType.STRING)
  private Status status;

  @OneToMany(mappedBy = "productAttribute")
  private List<ProductAttributeValue> productAttributeValues;
}
