package com.app.anyshop.cms.entity;

import jakarta.persistence.*;
import lombok.*;

import java.util.List;

@Getter
@Setter
@Builder
@Entity
@Table(name = "products")
@AllArgsConstructor
@NoArgsConstructor
public class Product extends Audit {
  @Id
  @GeneratedValue(strategy = GenerationType.UUID)
  private String id;

  @Column(unique = true, nullable = false)
  private String name;

  @Column(unique = true, nullable = false)
  private String slug;

  private String description;

  @Enumerated(EnumType.STRING)
  private Status status;

  private Double price;

  @OneToMany(mappedBy = "product")
  private List<ProductImages> images;

  @OneToMany(mappedBy = "product")
  private List<ProductAttributeValue> attributeValues;

  @ManyToOne(cascade = CascadeType.ALL)
  @JoinColumn(name = "cate_id")
  private ProductCategory productCategory;
}
