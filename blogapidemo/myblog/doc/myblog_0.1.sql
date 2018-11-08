-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema myblog
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema myblog
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `myblog` DEFAULT CHARACTER SET utf8mb4 ;
USE `myblog` ;

-- -----------------------------------------------------
-- Table `myblog`.`article`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `myblog`.`article` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(500) NULL,
  `content` LONGTEXT NULL,
  `create_time` DATETIME NULL,
  `update_time` DATETIME NULL,
  `type` INT NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `myblog`.`article_type`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `myblog`.`article_type` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `create_time` DATETIME NULL,
  `update_time` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `myblog`.`article_label`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `myblog`.`article_label` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(45) NULL,
  `type` INT NULL,
  `create_time` DATETIME NULL,
  `update_time` DATETIME NULL,
  PRIMARY KEY (`id`))
ENGINE = InnoDB;


-- -----------------------------------------------------
-- Table `myblog`.`article_re_label`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `myblog`.`article_re_label` (
  `re_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `label_id` INT NULL,
  `article_id` INT NULL,
  PRIMARY KEY (`re_id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
