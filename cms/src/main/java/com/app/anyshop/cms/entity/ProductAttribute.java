package com.app.anyshop.cms.entity;

import jakarta.persistence.*;
import lombok.*;

import java.util.List;

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
