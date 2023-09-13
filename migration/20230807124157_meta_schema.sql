-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE batch_job_instances
(
    id              BIGINT       NOT NULL AUTO_INCREMENT,
    job_key         VARCHAR(50)  NOT NULL,
    job_name        VARCHAR(100) NOT NULL,
    job_params      VARCHAR(512) NOT NULL,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP        NULL,
    deleted_at      TIMESTAMP        NULL,
    PRIMARY KEY (id),
    UNIQUE UNIQ_JOB_INST(job_key,job_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE batch_job_executions
(
    id               BIGINT        NOT NULL AUTO_INCREMENT,
    version          BIGINT        NOT NULL DEFAULT 0,
    job_instance_id  BIGINT        NOT NULL,
    job_name         VARCHAR(100)  NOT NULL,
    status           VARCHAR(10)   NOT NULL,
    exit_code        VARCHAR(2500) NOT NULL DEFAULT '',
    exit_message     VARCHAR(10000),
    created_at       TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at       TIMESTAMP         NULL,
    start_at         TIMESTAMP     NOT NULL DEFAULT '0000-00-00 00:00:00',
    end_at           TIMESTAMP     NOT NULL DEFAULT '0000-00-00 00:00:00',
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE batch_job_contexts
(
    id              BIGINT       NOT NULL AUTO_INCREMENT,
    job_instance_id BIGINT       NOT NULL,
    job_name        VARCHAR(100) NOT NULL,
    job_context     LONGTEXT,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP        NULL,
    deleted_at      TIMESTAMP        NULL,    
    PRIMARY KEY (id),
    UNIQUE UNIQ_JOB_INST(job_instance_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE batch_step_executions
(
    id                 BIGINT        NOT NULL AUTO_INCREMENT,
    version            BIGINT        NOT NULL DEFAULT 0,
    job_name           VARCHAR(100)  NOT NULL,
    job_instance_id    BIGINT        NOT NULL,
    job_execution_id   BIGINT        NOT NULL,
    step_name          VARCHAR(100)  NOT NULL,
    status             VARCHAR(10)   NOT NULL,
    commit_count       BIGINT,
    read_count         BIGINT,
    filter_count       BIGINT,
    write_count        BIGINT,
    read_skip_count    BIGINT,
    write_skip_count   BIGINT,
    process_skip_count BIGINT,
    rollback_count     BIGINT,
    execution_context  VARCHAR(2000) NOT NULL,
    step_context_id    BIGINT        NOT NULL,
    exit_code          VARCHAR(2500) NOT NULL DEFAULT '',
    exit_message       VARCHAR(10000),
    created_at         TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at         TIMESTAMP         NULL,
    start_at           TIMESTAMP     NOT NULL DEFAULT '0000-00-00 00:00:00',
    end_at             TIMESTAMP     NOT NULL DEFAULT '0000-00-00 00:00:00',
    PRIMARY KEY (id),
    KEY                IDX_JOB_EXECUTION_ID(job_execution_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd
-- +goose StatementBegin
CREATE TABLE batch_step_contexts
(
    id              BIGINT       NOT NULL AUTO_INCREMENT,
    job_instance_id BIGINT       NOT NULL,
    step_name       VARCHAR(100) NOT NULL,
    step_context    LONGTEXT,
    created_at      TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP        NULL,
    deleted_at      TIMESTAMP        NULL,    PRIMARY KEY (id),
    UNIQUE UNIQ_JOB_INST_STEP(job_instance_id, step_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
