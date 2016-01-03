package com.github.jacek99.snrteam.dao;

import javax.persistence.EntityManager;
import javax.persistence.EntityManagerFactory;
import javax.persistence.Persistence;

/**
 * Created by jfurmank on 1/2/16.
 */
public class Dao {

    private static EntityManagerFactory emf;

    public static void init() {
        emf = Persistence.createEntityManagerFactory("snrteam");
    }

    public static EntityManager getEntityManager() {
        return emf.createEntityManager();
    }

}
