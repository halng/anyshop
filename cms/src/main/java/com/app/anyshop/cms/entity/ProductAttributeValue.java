/*
 * *****************************************************************************************
 * Copyright 2024 By Hal Nguyen
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * *****************************************************************************************
 */

package com.app.anyshop.cms.entity;

import jakarta.persistence.*;
import lombok.*;

@Getter
@Setter
@Builder
@AllArgsConstructor
@NoArgsConstructor
@Entity
@Table(name = "product_attribute_value")
public class ProductAttributeValue extends Audit {
  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  private String id;

  private String value;

  @ManyToOne(cascade = CascadeType.ALL)
  @JoinColumn(name = "att_id")
  private ProductAttribute productAttribute;

  @ManyToOne(cascade = CascadeType.ALL)
  @JoinColumn(name = "prod_id")
  private Product product;
}
