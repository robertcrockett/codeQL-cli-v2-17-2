// Generated automatically from org.springframework.jdbc.core.JdbcOperations for testing purposes

package org.springframework.jdbc.core;

import java.util.Collection;
import java.util.List;
import java.util.Map;
import java.util.stream.Stream;
import org.springframework.jdbc.core.BatchPreparedStatementSetter;
import org.springframework.jdbc.core.CallableStatementCallback;
import org.springframework.jdbc.core.CallableStatementCreator;
import org.springframework.jdbc.core.ConnectionCallback;
import org.springframework.jdbc.core.ParameterizedPreparedStatementSetter;
import org.springframework.jdbc.core.PreparedStatementCallback;
import org.springframework.jdbc.core.PreparedStatementCreator;
import org.springframework.jdbc.core.PreparedStatementSetter;
import org.springframework.jdbc.core.ResultSetExtractor;
import org.springframework.jdbc.core.RowCallbackHandler;
import org.springframework.jdbc.core.RowMapper;
import org.springframework.jdbc.core.SqlParameter;
import org.springframework.jdbc.core.StatementCallback;
import org.springframework.jdbc.support.KeyHolder;
import org.springframework.jdbc.support.rowset.SqlRowSet;

public interface JdbcOperations
{
    <T> T execute(CallableStatementCreator p0, org.springframework.jdbc.core.CallableStatementCallback<T> p1);
    <T> T execute(PreparedStatementCreator p0, org.springframework.jdbc.core.PreparedStatementCallback<T> p1);
    <T> T execute(String p0, org.springframework.jdbc.core.CallableStatementCallback<T> p1);
    <T> T execute(String p0, org.springframework.jdbc.core.PreparedStatementCallback<T> p1);
    <T> T execute(org.springframework.jdbc.core.ConnectionCallback<T> p0);
    <T> T execute(org.springframework.jdbc.core.StatementCallback<T> p0);
    <T> T query(PreparedStatementCreator p0, org.springframework.jdbc.core.ResultSetExtractor<T> p1);
    <T> T query(String p0, Object[] p1, int[] p2, org.springframework.jdbc.core.ResultSetExtractor<T> p3);
    <T> T query(String p0, Object[] p1, org.springframework.jdbc.core.ResultSetExtractor<T> p2);
    <T> T query(String p0, PreparedStatementSetter p1, org.springframework.jdbc.core.ResultSetExtractor<T> p2);
    <T> T query(String p0, org.springframework.jdbc.core.ResultSetExtractor<T> p1);
    <T> T query(String p0, org.springframework.jdbc.core.ResultSetExtractor<T> p1, Object... p2);
    <T> T queryForObject(String p0, Object[] p1, int[] p2, java.lang.Class<T> p3);
    <T> T queryForObject(String p0, Object[] p1, int[] p2, org.springframework.jdbc.core.RowMapper<T> p3);
    <T> T queryForObject(String p0, Object[] p1, java.lang.Class<T> p2);
    <T> T queryForObject(String p0, Object[] p1, org.springframework.jdbc.core.RowMapper<T> p2);
    <T> T queryForObject(String p0, java.lang.Class<T> p1);
    <T> T queryForObject(String p0, java.lang.Class<T> p1, Object... p2);
    <T> T queryForObject(String p0, org.springframework.jdbc.core.RowMapper<T> p1);
    <T> T queryForObject(String p0, org.springframework.jdbc.core.RowMapper<T> p1, Object... p2);
    <T> int[][] batchUpdate(String p0, java.util.Collection<T> p1, int p2, org.springframework.jdbc.core.ParameterizedPreparedStatementSetter<T> p3);
    <T> java.util.List<T> query(PreparedStatementCreator p0, org.springframework.jdbc.core.RowMapper<T> p1);
    <T> java.util.List<T> query(String p0, Object[] p1, int[] p2, org.springframework.jdbc.core.RowMapper<T> p3);
    <T> java.util.List<T> query(String p0, Object[] p1, org.springframework.jdbc.core.RowMapper<T> p2);
    <T> java.util.List<T> query(String p0, PreparedStatementSetter p1, org.springframework.jdbc.core.RowMapper<T> p2);
    <T> java.util.List<T> query(String p0, org.springframework.jdbc.core.RowMapper<T> p1);
    <T> java.util.List<T> query(String p0, org.springframework.jdbc.core.RowMapper<T> p1, Object... p2);
    <T> java.util.List<T> queryForList(String p0, Object[] p1, int[] p2, java.lang.Class<T> p3);
    <T> java.util.List<T> queryForList(String p0, Object[] p1, java.lang.Class<T> p2);
    <T> java.util.List<T> queryForList(String p0, java.lang.Class<T> p1);
    <T> java.util.List<T> queryForList(String p0, java.lang.Class<T> p1, Object... p2);
    <T> java.util.stream.Stream<T> queryForStream(PreparedStatementCreator p0, org.springframework.jdbc.core.RowMapper<T> p1);
    <T> java.util.stream.Stream<T> queryForStream(String p0, PreparedStatementSetter p1, org.springframework.jdbc.core.RowMapper<T> p2);
    <T> java.util.stream.Stream<T> queryForStream(String p0, org.springframework.jdbc.core.RowMapper<T> p1);
    <T> java.util.stream.Stream<T> queryForStream(String p0, org.springframework.jdbc.core.RowMapper<T> p1, Object... p2);
    List<Map<String, Object>> queryForList(String p0);
    List<Map<String, Object>> queryForList(String p0, Object... p1);
    List<Map<String, Object>> queryForList(String p0, Object[] p1, int[] p2);
    Map<String, Object> call(CallableStatementCreator p0, List<SqlParameter> p1);
    Map<String, Object> queryForMap(String p0);
    Map<String, Object> queryForMap(String p0, Object... p1);
    Map<String, Object> queryForMap(String p0, Object[] p1, int[] p2);
    SqlRowSet queryForRowSet(String p0);
    SqlRowSet queryForRowSet(String p0, Object... p1);
    SqlRowSet queryForRowSet(String p0, Object[] p1, int[] p2);
    int update(PreparedStatementCreator p0);
    int update(PreparedStatementCreator p0, KeyHolder p1);
    int update(String p0);
    int update(String p0, Object... p1);
    int update(String p0, Object[] p1, int[] p2);
    int update(String p0, PreparedStatementSetter p1);
    int[] batchUpdate(String p0, BatchPreparedStatementSetter p1);
    int[] batchUpdate(String p0, List<Object[]> p1);
    int[] batchUpdate(String p0, List<Object[]> p1, int[] p2);
    int[] batchUpdate(String... p0);
    void execute(String p0);
    void query(PreparedStatementCreator p0, RowCallbackHandler p1);
    void query(String p0, Object[] p1, RowCallbackHandler p2);
    void query(String p0, Object[] p1, int[] p2, RowCallbackHandler p3);
    void query(String p0, PreparedStatementSetter p1, RowCallbackHandler p2);
    void query(String p0, RowCallbackHandler p1);
    void query(String p0, RowCallbackHandler p1, Object... p2);
}