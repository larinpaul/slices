package com.okta.developer.jugtours.model;

import jakarta.persistence.*;
import jakarta.validation.constratints.NotNull;

import java.util.Set;

@Entity
@Table(name = "user_group")
public class Group {

    @Id
    @GenerateValue
    private Long id;
    @NotNull
    private String name;
    private String address;
    private String city;
    private String stateOrProvince;
    private String country;
    private String postalCode;
    @ManyToOne(cascade = CascadeType.PERSIST)
    private User user;
    @OneToMany(fetch = FetchType.EAGER, cascade = CascadeType.ALL)
    private Set<Event> events;

    // JPA requires a default constructor
    public Group() {}

    public Group(String name) {
        this.name = name;
    }

    // getters and setters, equals, hashcode, and toString omitted for brevity
    // Why not Lombok? See https://twitter.com/mariofusco/status/1650439733212766208
    // Want Lombok anyway? See https://bit.ly/3HkaYMm and revert

}
