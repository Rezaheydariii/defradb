// Copyright 2022 Democratized Data Foundation
//
// Use of this software is governed by the Business Source License
// included in the file licenses/BSL.txt.
//
// As of the Change Date specified in that file, in accordance with
// the Business Source License, use of this software will be governed
// by the Apache License, Version 2.0, included in the file
// licenses/APL.txt.

package aggregates

import (
	"testing"

	testUtils "github.com/sourcenetwork/defradb/tests/integration/schema"
)

func TestSchemaAggregateSimpleCreatesUsersCount(t *testing.T) {
	test := testUtils.RequestTestCase{
		Schema: []string{
			`
				type users {}
			`,
		},
		IntrospectionRequest: `
			query IntrospectionQuery {
				__type (name: "users") {
					name
					fields {
						name
						args {
							name
							type {
								name
								inputFields {
									name
									type {
										name
									}
								}
							}
						}
					}
				}
			}
		`,
		ContainsData: map[string]any{
			"__type": map[string]any{
				"name": "users",
				"fields": []any{
					map[string]any{
						"name": "_count",
						"args": []any{
							map[string]any{
								"name": "_group",
								"type": map[string]any{
									"name": "users__CountSelector",
									"inputFields": []any{
										map[string]any{
											"name": "filter",
											"type": map[string]any{
												"name": "usersFilterArg",
											},
										},
										map[string]any{
											"name": "limit",
											"type": map[string]any{
												"name": "Int",
											},
										},
										map[string]any{
											"name": "offset",
											"type": map[string]any{
												"name": "Int",
											},
										},
									},
								},
							},
							map[string]any{
								"name": "_version",
								"type": map[string]any{
									"name": "users___version__CountSelector",
									"inputFields": []any{
										map[string]any{
											"name": "limit",
											"type": map[string]any{
												"name": "Int",
											},
										},
										map[string]any{
											"name": "offset",
											"type": map[string]any{
												"name": "Int",
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	testUtils.ExecuteRequestTestCase(t, test)
}

func TestSchemaAggregateSimpleCreatesUsersSum(t *testing.T) {
	test := testUtils.RequestTestCase{
		Schema: []string{
			`
				type users {}
			`,
		},
		IntrospectionRequest: `
			query IntrospectionQuery {
				__type (name: "users") {
					name
					fields {
						name
						args {
							name
							type {
								name
								inputFields {
									name
									type {
										name
										kind
										ofType {
											name
										}
									}
								}
							}
						}
					}
				}
			}
		`,
		ContainsData: map[string]any{
			"__type": map[string]any{
				"name": "users",
				"fields": []any{
					map[string]any{
						"name": "_sum",
						"args": []any{
							map[string]any{
								"name": "_group",
								"type": map[string]any{
									"name": "users__NumericSelector",
									"inputFields": []any{
										map[string]any{
											"name": "field",
											"type": map[string]any{
												"name": nil,
												"kind": "NON_NULL",
												"ofType": map[string]any{
													"name": "usersNumericFieldsArg",
												},
											},
										},
										map[string]any{
											"name": "filter",
											"type": map[string]any{
												"name":   "usersFilterArg",
												"kind":   "INPUT_OBJECT",
												"ofType": nil,
											},
										},
										map[string]any{
											"name": "limit",
											"type": map[string]any{
												"name":   "Int",
												"kind":   "SCALAR",
												"ofType": nil,
											},
										},
										map[string]any{
											"name": "offset",
											"type": map[string]any{
												"name":   "Int",
												"kind":   "SCALAR",
												"ofType": nil,
											},
										},
										map[string]any{
											"name": "order",
											"type": map[string]any{
												"name":   "usersOrderArg",
												"kind":   "INPUT_OBJECT",
												"ofType": nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	testUtils.ExecuteRequestTestCase(t, test)
}

func TestSchemaAggregateSimpleCreatesUsersAverage(t *testing.T) {
	test := testUtils.RequestTestCase{
		Schema: []string{
			`
				type users {}
			`,
		},
		IntrospectionRequest: `
			query IntrospectionQuery {
				__type (name: "users") {
					name
					fields {
						name
						args {
							name
							type {
								name
								inputFields {
									name
									type {
										name
										kind
										ofType {
											name
										}
									}
								}
							}
						}
					}
				}
			}
		`,
		ContainsData: map[string]any{
			"__type": map[string]any{
				"name": "users",
				"fields": []any{
					map[string]any{
						"name": "_avg",
						"args": []any{
							map[string]any{
								"name": "_group",
								"type": map[string]any{
									"name": "users__NumericSelector",
									"inputFields": []any{
										map[string]any{
											"name": "field",
											"type": map[string]any{
												"name": nil,
												"kind": "NON_NULL",
												"ofType": map[string]any{
													"name": "usersNumericFieldsArg",
												},
											},
										},
										map[string]any{
											"name": "filter",
											"type": map[string]any{
												"name":   "usersFilterArg",
												"kind":   "INPUT_OBJECT",
												"ofType": nil,
											},
										},
										map[string]any{
											"name": "limit",
											"type": map[string]any{
												"name":   "Int",
												"kind":   "SCALAR",
												"ofType": nil,
											},
										},
										map[string]any{
											"name": "offset",
											"type": map[string]any{
												"name":   "Int",
												"kind":   "SCALAR",
												"ofType": nil,
											},
										},
										map[string]any{
											"name": "order",
											"type": map[string]any{
												"name":   "usersOrderArg",
												"kind":   "INPUT_OBJECT",
												"ofType": nil,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	testUtils.ExecuteRequestTestCase(t, test)
}
